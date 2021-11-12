# GalGame论坛签到
目前支持的论坛有:
* [终点论坛](https://bbs.zdfx.net/)
    * 签到
    * 摇奖
* [花火学园](https://www.sayhuahuo.com/forum.php/)
    * 签到

## 使用方法

**在 GitHub Actions 中使用：**

1. Fork 本项目（顺便赏个 star 就更好了）
2. 前往 Actions 页面启用 GitHub Actions
3. 在`Secrets`中填写`ZDFX_COOKIE`和`HUAHUO_COOKIE`


**在本地使用：**

1. 安装 JRE
2. Clone 本项目
3. 在环境变量中添加与`Secrets`相同的值
4. /gradlew run