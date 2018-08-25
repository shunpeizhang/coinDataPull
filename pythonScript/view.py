# -*-coding=utf-8-*-


import pymysql
import matplotlib.pyplot as plt

import configInfo
import drawWave
import keyEventHandle


conn = pymysql.connect(host=configInfo.host,port=configInfo.port,user=configInfo.user,passwd=configInfo.passwd,db=configInfo.db)
fig = plt.figure()
keyEventHandle.Init(fig, conn)

drawWave.drawAllWave(conn, fig, 0)
plt.show()

















