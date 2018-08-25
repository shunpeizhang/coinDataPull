# -*-coding=utf-8-*-

import numpy as np
import talib as ta
import matplotlib
import matplotlib.pyplot as plt
import mpl_finance as mpf


#绘制bull线
def paintBull(ax, total, pos, totalDataLen, open, high, low, close, t):
    upperband, middleband, lowerband = ta.BBANDS(close)
    ax.plot(range(0, totalDataLen), upperband)
    ax.plot(range(0, totalDataLen), middleband)
    ax.plot(range(0, totalDataLen), lowerband)
    plt.legend(('upperband', 'middleband', 'lowerband'))


#绘制k线
def paintKLine(fig, total, pos, totalDataLen, open, high, low, close, t, isBull):
    ax = fig.add_subplot(total,1,pos)
    if isBull:
        paintBull(ax, total, pos, totalDataLen, open, high, low, close, t)

    pr = []
    for i in range(0,totalDataLen,1):
        pr.extend([[float(i),float(open[i]),float(close[i]),float(high[i]),float(low[i])]])

    mpf.candlestick_ohlc(ax, pr, width=0.8,colorup='r',colordown='g')
    ax.autoscale_view()
    plt.grid(False)
    plt.setp(plt.gca().get_xticklabels(), rotation=30)

#绘制macd
def paintMACD(fig, total, pos, totalDataLen, open, high, low, close, t):
    macd, signal, hist = ta.MACD(close, fastperiod=12, slowperiod=26, signalperiod=9)
    ax = fig.add_subplot(total,1,pos)
    ax.plot(range(0, totalDataLen), macd)
    ax.plot(range(0, totalDataLen), signal)
    ax.bar(range(0, 100), hist)
    plt.legend(('macd', 'signal', 'hist'))

#绘制rsi
def paintRSI(fig, total, pos, totalDataLen, open, high, low, close, t):
    rsi1 = ta.RSI(close, 5)
    rsi2 = ta.RSI(close, 10)
    rsi3 = ta.RSI(close, 14)

    ax = fig.add_subplot(total,1,pos)
    ax.plot(range(0, totalDataLen), rsi1)
    ax.plot(range(0, totalDataLen), rsi2)
    ax.plot(range(0, totalDataLen), rsi3)
    plt.legend(('5', '10', '14'))
