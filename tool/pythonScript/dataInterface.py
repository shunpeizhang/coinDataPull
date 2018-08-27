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
    #sql = "SELECT Open, High, Low, Close, ID FROM coin_data.`tb_coinData_60min` where coinType = 0 LIMIT %d, %d" % (start, count)
    #sql = "SELECT Open, High, Low, Close, ID FROM (SELECT ID, Amount, COUNT, OPEN, CLOSE, Low, High, Vol FROM coin_data.`tb_coinData_60min` WHERE ID <= 1531184400 AND coinType = 0 ORDER BY id DESC LIMIT 100) info ORDER BY ID"
    sql = "SELECT Open, High, Low, Close, ID FROM (SELECT ID, Amount, COUNT, OPEN, CLOSE, Low, High, Vol FROM coin_data.`tb_coinData_30min` WHERE ID <= 1531040400 AND coinType = 0 ORDER BY id DESC LIMIT 100) info ORDER BY ID"
    count = cur.execute(sql)
    for i in range(0, count):
        data = cur.fetchone()
        open.append(float(data[0]))
        high.append(float(data[1]))
        low.append(float(data[2]))
        close.append(float(data[3]))
        id.append(float(data[4]))

    #print(close)

    return np.array(open), np.array(high), np.array(low), np.array(close), np.array(id)


