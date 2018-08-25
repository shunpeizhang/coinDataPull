# -*-coding=utf-8-*-

from tkinter import *
import time
import matplotlib.pyplot as plt
import numpy as np

import drawWave
import configInfo

g_pos = 0
g_speed = 5

g_conf = None

def on_key_press(event):
    global g_pos
    global g_speed
    global g_conf
    if 'left' == event.key or 'right' == event.key:
        if 'left' == event.key:
            if g_pos >= g_speed:
                g_pos = g_pos - g_speed
        else:
            g_pos = g_pos + g_speed

        drawWave.drawAllWave(g_conf, g_pos)

    if  'up' == event.key or 'down' == event.key:
        if 'up' == event.key:
            g_speed = g_speed + 5
        elif 'down' == event.key:
            if g_speed > 5:
                g_speed = g_speed - 5

        print("speed:", g_speed)


def Init(conf):
    global g_conf

    g_conf = conf

    conf.g_fig_kline.canvas.mpl_disconnect(conf.g_fig_kline.canvas.manager.key_press_handler_id)#取消默认快捷键的注册
    conf.g_fig_kline.canvas.mpl_connect('key_press_event', on_key_press)



