In this video I'm going to show you how I approach setting up a Go Web application.
For persistence, we will be using a PostgreSQL database with a library called sqlx.

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
Now, in the command line, type `psql`, then paste this connection string and wrap it in double quotes. You should be able to connect to the database, just like before.

At long last, we can write some Go code. Make sure you are in the correct working directory, which is where you have initialized a Go module.
Create a file named main.go and open it in a code editor.
Define `package main` and `func main()`. Paste the connection URL and define it as a string constant.
We're going to have to import `sqlx` and `pq` in this file. If your code editor has automatic imports, it should handle importing `sqlx` for you, but you need to add the `pq` import yourself and prefix it with an underscore. This is because the package is not directly used in this file.
In the main function, connect to the database using `sqlx.MustConnect`. This function will panic if it doesn't manage to connect to the database.
Now, let's make a simple SQL query to check that everything works. Define a string variable called `version`.
Execute a SQL query and scan the result into the version variable.
If this method returned an error, log it to the standard output and exit. Otherwise, print the `version` variable.
If you run this project (git tag `step-1`), you should see the full version string of your Postgres installation.

Now, obviously, hard-coding the connection string inside the source code is not the best way going forward, so let's use environment variables instead.
In development, I like to use a tool called `direnv`.
`direnv` lets you define environment variables in a file called `.envrc`.
If you configure `direnv` correctly, it's going to evaluate the `.envrc` file every time you cd into the project directory and it will set your environment variables.
I won't be going into a detailed explanation on how to set up `direnv` on your machine, but you can read the friendly manual at https://direnv.net.

Create a file named `.envrc` in the working directory. In this file, we can define variables using standard shell syntax.
Let's export two variables. Set `PGDATABASE` to `goma_dev`. Then, set `DATABASE_URL` to the connection URL that we have previously hard-coded in the source code.
Save this file and open a new shell.
If you have configured `direnv` correctly, it should now display an error message, asking you to run `direnv allow`. If you haven't, you can still set the correct environment variables by running `source .envrc` in the shell.
You can check the value of `DATABASE_URL` by running `echo $DATABASE_URL` in the terminal.
The `PGDATABASE` variable sets the default database for Postgres, so if you type `psql` now, you should be connected directly to the database `goma_dev`.

In the `main.go` file, define a function called `MustGetEnv`, taking a string argument and returning a string.
Inside this function, read the value of an environment variable using `os.Getenv`. If the value is an empty string, log an error message and terminate the program.
Otherwise, return the value.
Now, replace the hardcoded connection string with a call to `MustGetEnv("DATABASE_URL")`. Since we are now using a function call, the value is no longer a stack-allocated constant, so we have to replace `const` with `var`.
When you run the program (git tag `step-2`), the output should remain the same.

This is already looking great, but we can do even better by creating a database table and writing some data into it.
In an application project, the best way to modify the database schema is using database migrations.
These are essentially SQL scripts checked into version control that a migration tool will execute in the correct order.
We're going to use a tool called `goose`. It's a powerful migration tool written in Go. On the website of goose, scroll down to this snippet, starting with `go install`, copy it and paste it in your terminal.
Then, create a directory to store migrations at `db/migrations`.
Now, let's configure goose using environment variables. Once again, open the file `.envrc` and add three new variables.
First, set the migration directory to `db/migrations`.
Then, set the migration driver to `postgres`.
Finally, set the `GOOSE_DBSTRING` to be equal to `DATABASE_URL`.
Open a new terminal. If you are using direnv, run `direnv allow`, otherwise source the `.envrc` file.
Now, create a migration using `goose create create_users sql`. Make sure to use a space between the name of the migration and "sql".
This command will create a new file in the `db/migrations` directory. Open that file in your code editor.
This file contains two sections, one named "up" for the changes we want to apply, and one named "down" for teardown instructions. The "down" section is optional.

In the "up" section, remove the default content and add this statement to install a database extension called `citext`.
This extension provides a data type for "case-insensitive text". We're going to use it for email addresses.
Then, we are going to create a table named "users".
We add an ID column, which is a 64-bit signed integer generated by the database, and will be the primary key for this table.
Add an email column, which is a "citext", or "case-insensitive text", cannot be null, and must be unique.
Then, a column named "password_hash", which is going to securely store the user's password.
Finally, add two columns named "inserted_at", and "updated_at". These columns represent the time when the record was created and updated. Both columns should be non-nullable, and both should default to the current UTC datetime.
In the "down" section, replace the default content with `drop table users;` and save the file.

Now, in the terminal, run `goose up`.
If you have set everything up correctly, the command should run your migration.
Connect to the database, type `\d users`, and press enter. As we can see, we now have a `users` table, and the columns are exactly the way we wanted.

When I started preparing this video two weeks ago, I was initially hoping to be able to show you much more than this.
However, this proved to be a much more daunting task than I could have imagined.
In the next videos, I want to cover password authentication using argon2, issuing access tokens signed with asymmetric keys, one-time passwords with TOTP, and WebAuthn, AKA passkeys.
If there is anything in particular you would like me to cover, please let me know in the comment section.
I hope you learned something, and thanks for watching.
