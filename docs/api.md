# API

### /connect

  `usi` と `isready` を渡し、 `readyok` が返るまで

### /close

  `quit`

### /option/list

  将棋エンジンのオプションの一覧を返す
  データ型は未決定

### /option/set

  `setoption name ~`

### /position/set

  POST
  `position sfen ~`

### /study/start

  `study` は研究的な意味。

### /study/stop

### /study/candidate/list

  エンジンの思考内容を取得

### /analyze/init

  棋譜解析前に棋譜を渡す

### /analyze/start

  解析開始
