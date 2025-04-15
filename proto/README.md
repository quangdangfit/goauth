# Proto

## 1. Overview
This project store all "proto" file all of the golang services.

## 2. Project layout
Project have 2 main folders: proto, api

Developer will edit or add new inside proto folder and script will generate api folder

## 3. Run command to install
Project require install python3
- Install python3
  ```brew install python3```

  ```python3 -m pip install --upgrade pip```

  ```pip3 install virtualenv```

  ```pip3 install GitPython```
## 4. Generate proto
  - Gen proto by folder: ```make gen path=auth```
  - Gen proto all folders: ```make gen_all```
  - OR you can install buf (brew install bufbuild/buf/buf - https://buf.build/docs/installation) and run command ```buf generate```