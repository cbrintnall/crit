# Crit

### Compilation

* clone repo
* `cd` into repo
* `make`
* `mv ./crit /usr/local/bin`

### Usage

#### Defining a .secrets file

`.secrets` files are just yaml. By default the `.secrets` file is looked for at `/home/{user}/.secrets`. To define a raw secret you'd use the following format:

```
secrets:
    - key: KEY
      value: VALUE
```

Environment variable `KEY` will be equal to `VALUE`. We can verify this by running:

```
crit start printenv
```

Which should output (amongst others):

```
KEY=VALUE
```

In the future, the goal is to support pulling from `Vault`, `AWS Secrets Manager`, and `GCP Secrets Manager`. You'll need to provide authentication to each service.

### Running

#### start

**to run:** `crit start <command>`

**example:** `crit start yarn start`

This will call `yarn start` and inject any secrets defined at `/home/{usr}/.secrets`.

Start is used to `exec` a program with injected variable. 