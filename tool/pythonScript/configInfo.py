# -*-coding=utf-8-*-

import matplotlib.pyplot as plt
import pymysql

host='47.97.202.23'
port = 3306
user='root'
passwd='root'
db ='coin_data'

handleDataLen = 100

class GConfInfo:
    g_fig_kline = None
    g_fig_macd = None
    g_fig_rsi = None

    g_conn = None


def Init(conf):
    conf.g_fig_kline = plt.figure()
    conf.g_fig_macd = conf.g_fig_kline
    conf.g_fig_rsi = conf.g_fig_kline

    conf.g_conn = pymysql.connect(host=host,port=port,user=user,passwd=passwd,db=db)




