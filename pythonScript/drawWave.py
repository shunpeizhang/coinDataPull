# -*-coding=utf-8-*-

import dataInterface
import configInfo
import arithmeticView



def drawAllWave(conf, dataStartPos):
    conf.g_fig_kline.clf()
    conf.g_fig_macd.clf()
    conf.g_fig_rsi.clf()

    open, high, low, close,t = dataInterface.getSomeData(conf.g_conn, dataStartPos, configInfo.handleDataLen)
    arithmeticView.paintKLine(conf.g_fig_kline, 3, 1, configInfo.handleDataLen, open, high, low, close, t, True)
    arithmeticView.paintMACD(conf.g_fig_macd, 3, 2, configInfo.handleDataLen, open, high, low, close, t)
    arithmeticView.paintRSI(conf.g_fig_rsi, 3, 3, configInfo.handleDataLen, open, high, low, close, t)

    conf.g_fig_kline.canvas.draw_idle()
    conf.g_fig_macd.canvas.draw_idle()
    conf.g_fig_rsi.canvas.draw_idle()


