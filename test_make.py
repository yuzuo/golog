# -*- coding: utf-8 -*-
"""
-------------------------------------------------
   File Name:   make_test_file
   Description: 测试文件
   Author:      shihaisheng
   date:        2021/3/16
-------------------------------------------------
   Change Activity:
                2021/3/16
-------------------------------------------------
# code is far away from bugs with the god animal protecting
    I love animals. They taste delicious.
              ┏┓      ┏┓
            ┏┛┻━━━┛┻┓
            ┃      ☃      ┃
            ┃  ┳┛  ┗┳  ┃
            ┃      ┻      ┃
            ┗━┓      ┏━┛
                ┃      ┗━━━┓
                ┃  神兽保佑    ┣┓
                ┃ 永无BUG！   ┏┛
                ┗┓┓┏━┳┓┏┛
                  ┃┫┫  ┃┫┫
                  ┗┻┛  ┗┻┛
"""
__author__ = 'shihaisheng'

with open("a.log", "a+") as f:
    for i in range(1, 20):
        print(i)
        f.write("""219.142.224.210 - - [15/Mar/2021:10:32:06 +0800] "GET http://test.vanwei.com.cn/ HTTP/1.1" 404 0 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36" "application/octet-stream" 0 Miss "C/404" Dynamic "-" 0.022 183.131.200.74
""")
    f.close()
