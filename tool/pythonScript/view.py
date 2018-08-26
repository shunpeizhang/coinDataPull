# -*-coding=utf-8-*-

import matplotlib.pyplot as plt

import configInfo
import drawWave
import keyEventHandle

conf = configInfo.GConfInfo()

#plt.xlim(0, 100)

configInfo.Init(conf)
keyEventHandle.Init(conf)

drawWave.drawAllWave(conf, 0)
plt.show()

















