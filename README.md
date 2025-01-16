# gerr
基于文章 https://juejin.cn/post/7182202759876182075  实现GO error的最佳实践！

# 项目背景
1. 在开发过程中，往往感觉现在错误处理不够优雅，总感觉error很难用，我要一层一层包出去，非常麻烦
    > e.g.
    > ```go
    > func main() {
    >     err := doSomething()
    >     if err != nil {
    >         // 由于需要定位err信息，
    >         //我们不得不每一次得到err都要输出和包装
    >         fmt.Println(err)
    >   }
    > }
    > func doSomething1() error {
    >   err := doSomething2()
    >   if err != nil {
    >       // 由于需要定位err信息，
    >       // 我们不得不每一次得到err都要输出和包装
    >       fmt.Println("doSomething1",err)  
    >       return errors.Wrap(err, "doSomething1")
    >   }
    > }
    > func doSomething2() error {
    >   // 由于需要定位err信息，
    >   // 我们不得不每一次得到err都要输出和包装
    >   err :=  errors.New("doSomething2")
    >   fmt.Println("doSomething2",err) 
    >   return err
    > }
    
2. 在浏览网站如何解决这个问题时，看到了[这篇文章]( https://juejin.cn/post/7182202759876182075),觉得博主写的文章非常好，但是很可惜没有看到文章源码，故心血来潮，根据博主的文章记录自己实现一个自己比较喜欢的错误处理包
    > e.g.
    > ```go
    > func main() {
    >     err := doSomething()
    >     if err != nil {
    >   
    >     }
    > }
    > func doSomething1() gerr.Error {
    >     err := doSomething2()
    >     if err != nil {
    >         return err  // 由于上层穿出来的类型是gerr.Error，所以上层可以不用进行log和记录堆栈信息，直接返回即可
    >     }
    > }
    > func doSomething2() gerr.Error {
    >     err :=  errors.New("doSomething2")
    >     if err != nil {
    >         err =  gerr.Wrap(err, "doSomething2") // 包装错误，这里同时会记录堆栈信息
    >         fmt.Println(err) // 仅在第一次抓到err的时候进行日志记录
    >         return err.(gerr.Error)
    >     }
    > }
# 待办事项

-[x] 定义好错误类型
-[x] 实现友好输出
-[x] 开发一些比较趁手的错误工具
-[ ] 更多的错误类型定义
