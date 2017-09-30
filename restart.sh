#!/bin/bash

killall beegoAutoDoc
chmod +u beegoAutoDoc
nohup ./beegoAutoDoc > "out.log" 2>&1 &
