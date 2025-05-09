git init -b main
这将在新仓库中直接创建名为 main 的初始分支，而不会创建 master 分支

git config --global init.defaultBranch main

如果您希望重新初始化当前目录下的 Git 仓库，可以直接运行：
git init
--initial-branch（或 -b）选项仅在初始化新的仓库时有效。如果仓库已经存在，Git 会忽略该选项，并重新初始化仓库结构，但不会更改当前的默认分支名称
如果您决定不再使用 Git，可以通过删除 .git 目录来完全移除版本控制：
rm -rf .git

git branch -m/M main将当前分支重命名为 main

