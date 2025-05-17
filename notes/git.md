git init -b main
这将在新仓库中直接创建名为 main 的初始分支，而不会创建 master 分支

git config --global init.defaultBranch main

如果您希望重新初始化当前目录下的 Git 仓库，可以直接运行：
git init
--initial-branch（或 -b）选项仅在初始化新的仓库时有效。如果仓库已经存在，Git 会忽略该选项，并重新初始化仓库结构，但不会更改当前的默认分支名称
如果您决定不再使用 Git，可以通过删除 .git 目录来完全移除版本控制：
rm -rf .git

git branch -m/M main将当前分支重命名为 main

在 Git 中，git push -u 中的 -u 是 --set-upstream 的简写。该选项的作用是将当前本地分支与远程分支建立关联（即设置“上游”分支）。这样，在后续的操作中，您可以简化命令，例如直接使用 git push 或 git pull，而无需每次都指定远程仓库和分支名。


git暂存区和工作区
工作区（Working Directory）：即你在本地计算机上看到的项目目录，包含所有实际的项目文件。
暂存区（Staging Area）：也称为索引（Index），用于暂时存放你准备提交的文件快照。
版本库（Repository）：存储项目历史版本的地方，位于 .git 目录下。
git add <file> 命令将修改添加到暂存区


git rm
作用：从 Git 的索引（即暂存区）和工作目录中删除指定的文件。
git rm --cached
作用：仅从 Git 的索引中删除指定的文件，保留工作目录中的文件。
git rm --cached <file> 命令的作用是将指定的文件从 Git 的索引中移除，即停止 Git 对该文件的跟踪


git restore  
-S, --staged          restore the index

git restore --staged
作用：将指定文件从暂存区（staging area）移除，使其恢复到上次提交的状态，但保留工作区中的修改。

git restore (--worktree) (default)
作用：将指定文件恢复到上次提交的状态，丢弃工作区中的修改。