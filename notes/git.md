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
