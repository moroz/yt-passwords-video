In this video I'm going to show you how I approach setting up a Go Web application.
For persistence, we will be using a PostgreSQL database with a library called sqlx.
In this video, we won't be touching the HTTP stack yet, but we will implement password authentication and JWTs.

Before we begin, make sure you have the latest Go toolchain installed.
As of this recording, the latest version is 1.22.0.
You will also need PostgreSQL. The latest version as of this recording is 16.2.
I will be developing this project using Debian 12 "Bookworm," but the experience should be more or less the same regardless of your operating system.

Create an empty directory named `goma` and cd into it.
"Goma" means "sesame" in Japanese.
I consider this word to be a good code name for a project related to password authentication in Go.
Not only does it begin with "go," the name of the programming language we are going to use, but it is also a reference to the story of "Ali Baba and the Forty Thieves," in which the phrase "open sesame" was used as a password to a treasure cave.

Initialize a Go module using the command `go mod init`, followed by a "module path."
This is going to be `github.com/`, your Github username, `/goma`.
My Github username is my surname, Moroz, so I initialize the project using `go mod init github.com/moroz/goma`.
Initialize a Git repository using `git init`.
Stage all changes in the working directory (`git add .`), and finally, create an initial commit (`git commit -m "Initial commit"`).

Now, let's try to connect to a database.
We're going to have to do two things: install some Go libraries, and write some Go code.
Let's start with the easy part: installing libraries.
The first library I'm going to install is sqlx, which is a wrapper over the standard library `database/sql` package.
Since we want to connect to a Postgres database, we also have to install `lib/pq`, which is a database driver for Postgres.
On the website of sqlx, I scroll down to this snippet and copy-and-paste it into the terminal.
Then, do the same thing for `lib/pq`.

Now, we're going to need to set up a database and define a connection string.
If you already know how to connect to a Postgres database using a connection URL, you can skip to the next section.

First, make sure that you can connect to Postgres using `psql`.
When you type `psql` in the command line, you should get connected to the default database, which is named the same as your system user.
In my case, it is my first name, karol.

If you have just installed Postgres for the first time, you can run into this error saying that the "role does not exist."
You can solve it by running the command shown on the screen (`sudo su postgres -c "createuser -s $(whoami)"`).
The first part of the command means that the command in quotes will be executed as the user "postgres".
The middle part means "create a Postgres superuser".
The last part is an expression that your shell should replace with your username.

If this command succeeds and you try to open a `psql` command line, you're going to run into another error, saying that your default database does not exist.
You can create the default database by simply calling the command `createdb`.
Then you should be able to open a Postgres console using `psql`.

Now let's set a password for the user "postgres".
In general, when connecting to a database on your own computer, the username and password don't really matter.
On the other hand, it is a common practice to use the default `postgres` user with the password `postgres`.
Still in the `psql` prompt, type the command shown on the screen `alter user postgres with password 'postgres';`.
Make sure to wrap the password in single quotes and to end the command with a semicolon (;).
If you see the message `ALTER ROLE`, it means that you have successfully changed the password.

Finally, let's create a database for the application.
Since the project is called "goma", I'm going to call the database "goma_dev", that's "dev" for "development".
The command to create a database is `createdb`, followed by the name of the database you want to create.
So I type `createdb goma_dev`. If there is no output, it means that the database has been successfully created.

In the terminal, we can very easily connect to this database by typing `psql goma_dev`, but if we want to make a connection using Go, we need to write a connection string.
On the screen, you can see a connection string in URL format.
This string specifies that we want to connect to the database server at `localhost`, using the username `postgres`, and the password `postgres`.
We want to connect to the database called `goma_dev`, and the part after the question mark means that we want to disable encryption, because we're not going to need it in development.
If you type `psql` in the command line, and then type this connection string, you should be able to connect to the database, just like before.
On some shells, you may have to wrap the URL in double quotes.

At long last, we can write some Go code. Make sure you are in the correct working directory, which is where you have initialized a Go module.
Create a file named main.go and open it in a code editor.
Define `package main` and `func main()`. Paste the connection URL and define it as a string constant.
We're going to have to import `sqlx` and `pq` in this file. If your code editor has automatic imports, it should handle importing `sqlx` for you, but you need to add the `pq` import yourself and prefix it with an underscore. This is because the package is not directly used in this file.
In the main function, connect to the database using `sqlx.MustConnect`. This function will panic if it doesn't manage to connect to the database.
Now, let's make a simple SQL query to check that everything works. Define a string variable called `version`.
Using the `db.QueryRow` method, execute a SQL query and read the result into the `version` variable, passing a reference to the `.Scan` method.
If this method returned an error, log it to the standard output and exit. Otherwise, print the `version` variable.
If you run this project (git tag `step-1`), you should see the full version string of your Postgres installation.

Now, obviously, hard-coding the connection string inside the source code is not the best way going forward, so let's move it out to environment variables.
In development, I like to use a tool called `direnv` to manage my environment variables.
`direnv` configures environment variables based on a configuration file called `.envrc`.
If you configure `direnv` correctly, it's going to evaluate the `.envrc` file every time you cd into the project directory and it will set your environment variables.
I won't be going into a detailed explanation on how to set up `direnv` on your machine, but you can read the friendly manual at https://direnv.net.

Create a file named `.envrc` in the working directory. Inside this file, type `export DATABASE_URL=`, and then paste the connection URL.
If you have configured `direnv` correctly, it should now display an error message, asking you to run `direnv allow`. If you haven't, you can still set the correct environment variables by running `source .envrc` in the shell.
You can check the value of `DATABASE_URL` by running `echo $DATABASE_URL` in the terminal. Remember to prefix the variable name with a dollar sign.

In the `main.go` file, define a function called `MustGetEnv`, taking a string argument and returning a string.
Read the value of an environment variable using `os.Getenv`. If the value is an empty string, log an error message and terminate the program.
Otherwise, return the value.
Now, replace the hardcoded connection string with a call to `MustGetEnv("DATABASE_URL")`. Since we are now using a function call, the value is no longer a stack-allocated constant, so we have to replace `const` with `var`.
When you run the program (git tag `step-2`), the output should remain the same.

This is already looking great, but we can do even better by creating a database table and writing some data into it.
In an application project, the best way to modify the database schema is using database migrations.
These are essentially SQL scripts checked into version control that a migration tool will execute on the database in the correct order.
We can use a tool called `golang-migrate` CLI. It's kind of a pain to install, but on Linux and Mac you can install it using Homebrew, by typing `brew install golang-migrate`. Other installation options are described on this website.
Create a directory to store migrations at `db/migrations`.
The command to generate a migration is shown on the screen. This instructs `golang-migrate` to create a migration script, with the `.sql` file extension, inside the `db/migrations` directory, and we're going to call it "create_users".
If you execute this command, you should get two files, one ending with `.up.sql` and one ending with `.down.sql`.
When you apply a migration script, the "up" script is executed, and when you revert one, the "down" script is executed.

Open the "up" migration in your text editor.
This statement will install a database extension called "citext", or "case-insensitive text".
"Case-insensitive" means that capital letters are treated the same as small letters, just like in email addresses.
Then, we are going to create a table named "users". We add an ID column, which is a 64-bit signed integer generated by the database, and will be the primary key for this table.
Add an email column, which is a "citext", or "case-insensitive text", cannot be null, and must be unique.
Then, a column named "password_hash", which is going to securely store the user's password.
Finally, add two columns named "inserted_at", and "updated_at". These columns represent the time when the record was created and updated. Both columns should be non-nullable, and both should default to the current datetime in the UTC timezone.
Remember to double-check that everything in this file is syntactically correct, because the `golang-migrate` tool is actually quite primitive, and if there are any errors, it will lock the database and you will have to fix it manually.
Now, open the "down" migration file. Here, just type `drop table users;` and save. There is no need to revert the first statement in the "up" migration as it doesn't really influence the overall database structure.

Now, the command to apply or revert the migrations is quite complicated, and I don't want to have to type it every time, so I'm going to create a file called "Makefile", with a capital M in front and all the remaining letters are small.
This is a configuration file for a tool called `make`, and it is an easy way to define common commands.
In this file, I'm first going to define a task named `check-env`.
Using the `test` command, I check if the value of `DATABASE_URL` is non-empty.
Then, I add the task `db.migrate`, which depends on `check-env`.
This task runs the 

