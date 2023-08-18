That is 4 test.

Git GO

changge 4 test

Git init创建

add

commit -m \<message\>

status 查看状态（如改动信息）

diff 查看差异

log 查看历史

HEAD 当前版本

git reset --hard HEAD^^ 回退版本

reflog  记录每一次命令，可以回到未来的版本

checkout -- file可以丢弃工作区的修改，让这个文件回到最近一次git commit或git add时的状态

reset HEAD file可以把暂存区的修改撤销掉（unstage），重新放回工作区：

git rm file 删除文件

git push  把当前分支master推送到远程。 -u参数，Git不但会把本地的master分支内容推送的远程新的master分支，还会把本地的master分支和远程的master分支关联起来

分支

查看分支：git branch

创建分支：git branch name

切换分支：git checkout name或者git switch name

创建+切换分支：git checkout -b name或者git switch -c name

合并某分支到当前分支：git merge name

删除分支：git branch -d name