# 从GitHub Pages到Container

第一版网站启动的足够简单。我在电脑上用`vscode`写笔记，用3行[`bash`脚本](https://github.com/lzztt/onenow/blob/b7cdde585b23c219adfe169bde28b5d9cb232d59/update_home.sh)更新笔记列表，用`git`发布到GitHub Pages，就完成了。

现在需要把网站从GitHub Pages搬到Container。这样就可以把网站打包，在不同环境（开发电脑，服务器）中运行，从而实现“运行环境”自由。

昨晚折腾了一晚上，发现Jekyll最新的稳定版（4.2.2）的Docker image居然是坏的。[一堆问题](https://github.com/envygeeks/jekyll-docker/issues)半年了没人修。

今天换到一年前的4.2.0版本就可以了。看来Jekyll这个Ruby项目的维护堪忧啊。

这也体现了container的优势：整个运行环境打包成一个image，无需安装，不依赖（污染）系统。再也不用担心旧版本不能运行和升级后的新版本崩溃了。
> 它崩由它崩，清风拂山岗；
> 它旧由它旧，明月照大江。

即使Jekyll现在的维护不给力，旧版image依然可以流畅的运行在不同系统上。
