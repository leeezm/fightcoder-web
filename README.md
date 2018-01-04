####项目名称:
#分布式在线编程平台
####项目功能:
    在线编程平台(Online Judge)，用户可以做题并对自己代码的正确性进行验证的一个平台。
    
    网站最核心的功能：
    (1)用户在线提交代码，系统将结果返回给用户。
    (2)网站支持比赛功能，多个用户进行在线答题，系统给出评测并对结果进行排名。
    
    相关的其他功能：
    (1)为方便用户，提供第三方登录注册以及题目分享接口；
    (2)网站根据用户的AC题目数量，比赛情况等信息，给出相应的排名信息；
    (3)提供用户的提交记录等信息，并且提供收藏夹，便于用户记录感兴趣的题目；
    (4)提供用户交流的论坛功能；
    (5)提供用户进行题目上传，测试数据报错等功能；
    (6)提供站内信以及系统通知，便于用户交流以及通知事情。

####项目难点
    (1)网站在用户访问量很大的时候如何保证低延迟的用户编程体验；
    (2)如何实现网站的分布式管理；
    (3)如何使用缓存来对网站的一些功能进行优化；
    (4)如何保证该网站平台的安全性,以及用户信息的安全性。

####项目技术:
    (1)后台采用Gin框架，数据库采用MySQL数据库，后台向前端提供访问接口；
    (2)采用redis缓存；
     未完，待续...