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
Stage all changes in the working directory using `git add .`, and finally, create an initial commit using `git commit -m "Initial commit"`.

Now, let's try to connect to a database.
In order to be able to connect to a database, we need to do two things: we need to define a connection string, and we need to install some Go libraries.
Let's start with the easy part: installing libraries.
I'm going to install sqlx, which is a wrapper over the Go standard library `database/sql` package.
On the website of sqlx, I scroll down to this snippet and copy-and-paste it in the terminal.
Then, install `lib/pq`, which is a database driver for Postgres.

Now, we're going to need to define a connection string.
If you know how to connect to a Postgres database using a connection string, you can skip to the next section.

First, make sure that you can actually connect to Postgres using `psql`.
When you type `psql` in the command line, you should get connected to the default database, which is named the same as your system user.

If you have just installed Postgres for the first time, you can run into an error saying that the "role does not exist."
You can solve it by running this command.
What this does is it runs a command as the `postgres` system user to create a database role.
The `-s` flag means that you want to create a superuser.
`$(whoami)` will be replaced with your username.

If this command succeeds and you try to open a `psql` command line, you're going to run into another error, saying that the database does not exist.
You can create the default database using the command `createdb`.
Then you should be able to open a Postgres console using `psql`.
Now let's set a password for the user "postgres".
Still in the `psql` prompt, run `alter user postgres with password 'postgres';`.
Make sure to end your command with a semicolon.
If it says `ALTER ROLE`, it means that the command was successful.

Now, let's create a database for the application.
Since the project is called "goma", I'm going to call the development database "goma_dev".
Creating the database is as simple as running `createdb goma_dev` in the terminal.
Now, we can try to connect to the database with a connection string by typing `psql`, double quote, `postgres`, colon, slash slash, `postgres`, colon, `postgres@localhost`, slash, `goma_dev`, question mark, `sslmode` equals `disable`, double quote.
This means that we want to connect as user `postgres`, with password `postgres`, to the server running on `localhost`, to the database `goma_dev`.
The last part means that we want to disable encryption, because we don't need it when we connect to a database on the same machine.

We could hard-code the connection string in the source files, but a better way to set it is using environment variables.
In development, I use a command-line tool called `direnv` to manage my environment.
`direnv` configures environment variables based on a configuration file called `.envrc`.
If you configure `direnv` correctly, it's going to evaluate the `.envrc` file every time you cd into the project directory and it will set your environment variables.
I won't be going into a detailed explanation on how to set up `direnv` on your machine, but you can read the friendly manual at https://direnv.net.
