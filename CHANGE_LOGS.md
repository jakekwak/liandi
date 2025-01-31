## v1.1.8 / 2020-08-29

### 引入特性

* [内容块被引用数统计](https://github.com/88250/liandi/issues/138)
* [导出为标准 Markdown](https://github.com/88250/liandi/issues/192)

### 改进功能

* [关系图优化](https://github.com/88250/liandi/issues/190)
* [删除、重命名文档以及关闭笔记本时关系图、编辑器、大纲页签的同步](https://github.com/88250/liandi/issues/193)
* [初始化时根据设备 Locale 设置语言](https://github.com/88250/liandi/issues/194)

### 修复缺陷

* [标题和有序列表导入的问题](https://github.com/88250/liandi/issues/191)

## v1.1.7 / 2020-08-28

### 引入特性

* [支持指定配置目录路径](https://github.com/88250/liandi/issues/78)

### 改进功能

* [稳定关系图布局](https://github.com/88250/liandi/issues/180)
* [关系图去掉根块子块](https://github.com/88250/liandi/issues/182)
* [跳转块定义定位优化](https://github.com/88250/liandi/issues/184)

### 修复缺陷

* [大纲上下分割问题](https://github.com/88250/liandi/issues/181)
* [导入异常引起内核崩溃导致的白屏问题](https://github.com/88250/liandi/issues/183)
* [新建空文件名文档问题](https://github.com/88250/liandi/issues/185)

## v1.1.6 / 2020-08-26

### 引入特性

* [实现引力布局](https://github.com/88250/liandi/issues/179)

### 改进功能

* [创建文档快捷入口](https://github.com/88250/liandi/issues/140)
* [大纲独立页签](https://github.com/88250/liandi/issues/146)
* [关系图增加选中操作](https://github.com/88250/liandi/issues/149)
* [编辑器改进](https://github.com/88250/liandi/issues/167)
* [选中内容块后关系图上对应的节点高亮](https://github.com/88250/liandi/issues/168)
* [优化全局搜索结果渲染](https://github.com/88250/liandi/issues/170)
* [简化关系图去掉“关联块”](https://github.com/88250/liandi/issues/175)
* [编辑器页签加入复制 ID](https://github.com/88250/liandi/issues/177)
* [优化关系图视觉效果](https://github.com/88250/liandi/issues/178)

### 修复缺陷

* [数学块公式块不能展现和引用问题](https://github.com/88250/liandi/issues/165)
* [图片首次加载渲染问题](https://github.com/88250/liandi/issues/169)

## v1.1.5 / 2020-08-25

### 引入特性

* [内容块复制 ID](https://github.com/88250/liandi/issues/145)
* [为内容块增加显示标识](https://github.com/88250/liandi/issues/158)

### 改进功能

* [点击帮助后自动展开用户指南笔记本](https://github.com/88250/liandi/issues/93)
* [已有文件打开应激活](https://github.com/88250/liandi/issues/148)
* [编辑后文档自动保存](https://github.com/88250/liandi/issues/160)

### 文档相关

* [增加 Markdown 完整示例文档](https://github.com/88250/liandi/issues/161)

### 修复缺陷

* [mac window-all-closed 键盘监听报错](https://github.com/88250/liandi/issues/162)

## v1.1.4 / 2020-08-24

* 完全重写
* 支持块级引用和双向链接

## v1.1.3 / 2020-07-30

### 引入特性

* [分屏预览模式支持语法高亮](https://github.com/88250/liandi/issues/77)

### 改进功能

* [重写编辑器分屏预览模式](https://github.com/88250/liandi/issues/62)

### 开发重构

* [去除 Setext 标题解析开关](https://github.com/88250/liandi/issues/86)

## v1.1.2 / 2020-06-28

### 改进功能

* [Markdown 中文排版段首缩进配置开关](https://github.com/88250/liandi/issues/73)

### 修复缺陷

* [编辑器插入文件时的检查问题](https://github.com/88250/liandi/issues/67)
* [链滴中点击大纲中的标题，编辑器正文和预览无法跳转到相应位置](https://github.com/88250/liandi/issues/68)

## v1.1.1 / 2020-05-21

### 改进功能

* [是否默认展示大纲](https://github.com/88250/liandi/issues/57)
* [Add language config option in the first screen](https://github.com/88250/liandi/issues/58)

### 修复缺陷

* [修复 XSS 漏洞](https://github.com/88250/liandi/issues/56)

## v1.1.0 / 2020-05-04

### 引入特性

* [支持导出 PDF、HTML 和微信公众号编辑器](https://github.com/88250/liandi/issues/52)

### 修复缺陷

* [编辑器即时渲染、分屏预览模式粘贴为纯文本失效](https://github.com/88250/liandi/issues/50)
* [搜索框切换 bug](https://github.com/88250/liandi/issues/51)

## v1.0.0 / 2020-04-30

### 引入特性

* [支持类似 Typora 的即时渲染模式（保留 Markdown 标记符）](https://github.com/88250/liandi/issues/10)

### 改进功能

* [Double Shift 快捷键容易误触](https://github.com/88250/liandi/issues/45)
* [增加 Setext 标题解析开关](https://github.com/88250/liandi/issues/46)
* [光标位置保留](https://github.com/88250/liandi/issues/49)

### 文档相关

* [Add README in English](https://github.com/88250/liandi/issues/48)

### 修复缺陷

* [自动拉取图片报错](https://github.com/88250/liandi/issues/47)

## v0.1.4 / 2020-04-07

### 引入特性

* [可配置是否自动拉取图片到本地](https://github.com/88250/liandi/issues/39)

### 改进功能

* [所见即所得模式下链接为空时会变为 http://127.0.0.1:6807/webdav/xxx](https://github.com/88250/liandi/issues/40)
* [复制为纯文本](https://github.com/88250/liandi/issues/41)

### 修复缺陷

* [文件名中 + 会变为空格](https://github.com/88250/liandi/issues/42)
* [带 HTML 标签的文件名打不开](https://github.com/88250/liandi/issues/43)
* [选中文字后 ctrl+shift+v，选中的文字没有被删除](https://github.com/88250/liandi/issues/44)

## v0.1.3 / 2020-03-10

### 引入特性

* [支持 Graphviz](https://github.com/88250/liandi/issues/35)

### 改进功能

* [支持隐藏编辑器工具栏](https://github.com/88250/liandi/issues/18)
* [打开选择不挂载，点击关闭按钮失效](https://github.com/88250/liandi/issues/33)
* [支持本地绝对路径渲染图片](https://github.com/88250/liandi/issues/34)
* [支持 ToC](https://github.com/88250/liandi/issues/37)
* [支持脚注](https://github.com/88250/liandi/issues/38)

## v0.1.2 / 2020-02-25

### 改进功能

* [文件名移除 .md](https://github.com/88250/liandi/issues/28)

### 开发重构

* [获取内核进程调试信息](https://github.com/88250/liandi/issues/29)

### 修复缺陷

* [无响应问题](https://github.com/88250/liandi/issues/30)
* [下载太慢，使用自建服务器](https://github.com/88250/liandi/issues/31)

## v0.1.1 / 2020-02-23

### 改进功能

* [更新提示改进](https://github.com/88250/liandi/issues/23)
* [新建文档后选中并打开](https://github.com/88250/liandi/issues/24)
* [macOS 应用菜单细化](https://github.com/88250/liandi/issues/26)
* [多主题改进](https://github.com/88250/liandi/issues/27)

### 修复缺陷

* [工具提示被遮挡](https://github.com/88250/liandi/issues/22)

## v0.1.0 / 2020-02-21

实现笔记应用基础功能，第一次公开发布。
