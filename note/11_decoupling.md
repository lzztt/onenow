# 模块解耦与演化

网站写到了[第四版](https://github.com/lzztt/onenow/releases/tag/v4.0.0)，进行了前端、后端和数据的分离。

分离的好处是解耦，各模块可以独立快速的扩展和演进。

- [前端](https://github.com/lzztt/onenow/tree/248f98db626392a15955b6542250a99444fc5964/frontend/src)：150行`TypeScript`和`React`代码，负责用户界面。
- [API](https://github.com/lzztt/onenow/tree/248f98db626392a15955b6542250a99444fc5964/proto)：17行`gRPC`代码，标准化各模块间的数据操作接口。
- [后端](https://github.com/lzztt/onenow/tree/248f98db626392a15955b6542250a99444fc5964/backend)：96行`Go`代码，负责进行数据操作。
- [数据](https://github.com/lzztt/onenow/tree/248f98db626392a15955b6542250a99444fc5964/note)：10个note，200行`Markdown`文字，负责存储数据。
