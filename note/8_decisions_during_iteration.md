# 迭代中的取舍

网站到第三版了，用了`React`和`TypeScript`，从此进入正式的前端开发。

目前网站还是静态：内容的更新需要重新编译静态文件和发布。

## 前三版本的取舍

- [第一版](https://github.com/lzztt/onenow/tree/b7cdde585b23c219adfe169bde28b5d9cb232d59)：`GitHub Pages`，选择“简单”和“速度”（第一天上线）。
- [第二版](https://github.com/lzztt/onenow/tree/657ee059bb3822bdfbb12bb402c86a77dc90890a)：`Container` + `Jekyll`，选择“简单”和“灵活性”（运行环境可迁移性），这样网站可以运行在自己的开发电脑和网站服务器上。
- [第三版](https://github.com/lzztt/onenow/tree/354722daaf988dd47f35ebf4fdd72c1f95cec164)：`React` + `TypeScript`，选择“生态兼容性”和“灵活性”，这样可以接入`React`和`JavaScript`的庞大生态，开始主流的前端开发。

“简单”一直是选择的侧重点，而不是“性能”或者“新颖”。如果一个选项经过一天的测试，还存在未解决的问题，而不能令人满意，则它会因为“不简单”而被放弃。


## 第三方组件/工具的选择
引入第三方组件/工具是为了方便好用，解决问题，使精力专注。而不是让它们不断的制造问题。

选择标准：
- 使用量、稳定性
- 社区和生态的大小
- 兼容性、整合能力

具体选择：
- `React`，而不是其他前端库。
- `BootStrap`，而不是其他样式库。
- `TypeScript`，而不是`JavaScript`（或其他前端语言）。可以借助IDE（vscode）的类型检查，在编码最早期避免类型方面的bug。
- `yarn v1`，而不是`v3`（或`npm`）。`yarn v1`兼具速度和稳定性。`yarn v3`运行`Create React App`生成的初始代码时有错误。
- `Create React App`，而不是`Vite`。因为`Vite`需要配置。我花了一晚上时间配置和测试，最后jest运行测试有问题。`Create React App`生成的初始代码从工程角度来看更成熟和易维护。
