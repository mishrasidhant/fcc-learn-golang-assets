# Assets for "Learn Go" on FreeCodeCamp

This is a snapshot of the code samples for the ["Learn Go" course](https://boot.dev/courses/learn-golang) on [Boot.dev](https://boot.dev) at the time the video for FreeCodeCamp was released on YouTube. If you want the most up-to-date version of the code, please visit the official [Boot.dev course](https://boot.dev/courses/learn-golang). Otherwise, if you're looking for the files used in the video, you're in the right place!

* [Course code samples](/course)
* [Project steps](/project)

## License

You are free to use this content and code for personal education purposes. However, you are *not* authorized to publish this content or code elsewhere, whether for commercial purposes or not. 

## Forked Branch workflow
### Remotes
- `origin` points to forked repo
- `upstream` points to repository that this repo was forked from

```
git remote rename origin upstream
git remote add origin git@github.com:mishrasidhant/fcc-learn-golang-assets.git
```
### Workflow
- `main` branch will always be up to date with upstream
- `forked` branch will be tracked as my "main" branch where all changes will live

#### Keeping main updated
```
# Create main branch
git checkout --track origin/main
# Checkout to main branch (if already created)
git checkout main
# Sync main branch with upstream
git pull upstream
```
#### Work on forked branch
```
git checkout forked
# Make changes
# Commit and push changes to forked directly
```