#!/bin/bash

d=$(cd $(dirname $0); pwd)
conf=$d/app.yml

if [[ ! -f $conf ]]; then
  echo "app.yml が存在しません。"
  echo "設定ファイルを用意してください。"
  echo "場所: $conf"
  exit 1
fi

sbs=$d/$(ls -1 $d | grep sbserver)

if [[ ! -x $sbs ]]; then
  echo "実行権限がありません。 ファイル: $sbs"
  exit 1
fi

echo $sbs
$sbs -app_config $conf
