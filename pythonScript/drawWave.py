# -*-coding=utf-8-*-

import dataInterface
import configInfo
import arithmeticView


def drawAllWave(conn, fig, dataStartPos):
    fig.clf()

    open, high, low, close,t = dataInterface.getSomeData(conn, dataStartPos, configInfo.handleDataLen)
    arithmeticView.paintKLine(fig, 3, 1, configInfo.handleDataLen, open, high, low, close, t, True)
    arithmeticView.paintMACD(fig, 3, 2, configInfo.handleDataLen, open, high, low, close, t)
    arithmeticView.paintRSI(fig, 3, 3, configInfo.handleDataLen, open, high, low, close, t)

    fig.canvas.draw_idle()



