# Branching System

The NeplodeAPI uses a **Dev-Staging-Prod** branching flow.

## Dev Branch

The dev branch is where all in-development commits go. If you just finished a new system but haven't done full bug-checking or something along those lines. It goes to dev.

## Staging

Once the code is complete and checked to be free of any API-Breaking bugs, it is commited to staging. Staging is where all complete features or commits are stored.

## Prod

When we are ready to release an update, all commits from staging are commited to prod, which is then deployed to the API on our server.
