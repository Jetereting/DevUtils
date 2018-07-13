#!/bin/bash

killall DevUtils
chmod +u DevUtils
nohup ./DevUtils > "out.log" 2>&1 &
