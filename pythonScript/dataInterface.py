# -*-coding=utf-8-*-

import numpy as np

#得到一定数量的open, high, low, close, id
def getSomeData(mysqlConn, start, count):
    open = []
    high = []
    low = []
    close = []
    id = []

    cur = mysqlConn.cursor()
    sql = "SELECT Open, High, Low, Close, ID FROM coin_data.`tb_coinData_60min` where coinType = 0 LIMIT %d, %d" % (start, count)
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


