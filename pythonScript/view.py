# -*-coding=utf-8-*-


import numpy as np
import talib as ta
#import MySQLdb
import pymysql
import matplotlib
import matplotlib.pyplot as plt
import mpl_finance as mpf


host='47.97.202.23'
port = 3306
user='root'
passwd='root'
db ='coin_data'
conn = pymysql.connect(host=host,port=port,user=user,passwd=passwd,db=db)



#得到一定数量的open, high, low, close
def getSomeData(mysqlConn, count):
    open = []
    high = []
    low = []
    close = []
    id = []

    cur = mysqlConn.cursor()
    sql = "SELECT Open, High, Low, Close, ID FROM coin_data.`tb_coinData_15min` where coinType = 0 LIMIT %d" % (count)
    count = cur.execute(sql)
    for i in range(0, count):
        data = cur.fetchone()
        open.append(float(data[0]))
        high.append(float(data[1]))
        low.append(float(data[2]))
        close.append(float(data[3]))
        id.append(float(data[4]))
        #print(data)

    return np.array(open), np.array(high), np.array(low), np.array(close), np.array(id)

open, high, low, close,t = getSomeData(conn, 100)
print(open, high, low, close)


pr = []
for i in range(0,100,1):
    pr.extend([[
        float(i)
        ,float(open[i])
        ,float(close[i])
        ,float(high[i])
        ,float(low[i])]]
        )






macd, signal, hist = ta.MACD(close, fastperiod=12, slowperiod=26, signalperiod=9)

fig = plt.figure()
ax = fig.add_subplot(2,1,1)
#ax.plot(range(0, 100), close)
# ax.plot(range(0, 100), macd)
# ax.plot(range(0, 100), signal)
# ax.plot(range(0, 100), hist)

mpf.candlestick_ohlc(ax,pr, width=0.4,colorup='r',colordown='g')
plt.grid(False)
# ax.xaxis_date()
ax.autoscale_view()
plt.setp(plt.gca().get_xticklabels(), rotation=30)

ax1 = fig.add_subplot(2,1,2)
#ax.plot(range(0, 100), close)
ax1.plot(range(0, 100), macd)
ax1.plot(range(0, 100), signal)
#ax1.plot(range(0, 100), hist)
ax1.bar(range(0, 100), hist)
plt.legend(('macd', 'signal', 'hist'))

plt.show()



close = np.random.random(100)
output = ta.CDL2CROWS(open, high, low, close)



print(output)









