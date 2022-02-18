# Stragegy (策略模式)
定義一系統的演算法，把它們一個個封裝起來，並且使它們可以相互替換。Strategy模式使演算法可以獨立於使用它們的客戶而變化

## 理解
在一個既存的系統中，因應任務所需，必需要引入第三方的函式庫，但函式庫中提供的介面與既存系統無法銜接，這個時候可用Adapter來轉接第三方函式庫，使既存系統可以為第三方函式庫順利合作。

## 定義
* 意圖：使控制範圍之外的一個原有物件與某個介面匹配。
* 問題：系統的資料和行為都正確，但是介面不符，通常用於必須從抽象類別衍生類別時。
* 方案：Adapter模式提供了具有所需介面的包裝類別(wrapper)。
* 角色：Adapter改變了Adapee的介面，使Adapee與Adapter的基礎類別Target匹配。如此一起使用者就可以使用Adapee了，就好像它是Target類型。
* 效果：Adapter模式使原有物件能夠適應新的類別結構，不受其介面限制。
* 實作：將原有類別包含在另一個類別之中，讓包含類別與需要的介面匹配，呼叫被複合類別的方法。

## 變形
* 物件Adapter模式：相依於一個轉換物件來包含另一個被轉換物件
* 類別Adapter模式：透過多重繼承來建立一個新的類別，該類別同時從兩個類別繼承
  * 定義其介面的抽象類別(公開繼承)
  * 存取其實作的原有類別(私有繼承)
每個被包裝的方法都呼叫其對應的私有繼承方法

## Adapter模式UML
![image](https://github.com/Lornzo/DesignPattern/blob/main/Adapter/images/pattern.png)

## Adapter模式 VS Facade模式
* Facade模式簡化了介面
* Adapter模式將一個已有的介面轉換成另一個介面  
![image](https://github.com/Lornzo/DesignPattern/blob/main/Adapter/images/compare.png)

## 範例說明
* 有一個Function「Client」，會比較數字10是不是大於某個數字
* 有一個第三方的函式庫，其中有一個Function能夠以字串格式回傳一個隨機的數字
* 當我們不想，或是沒有辦法重新寫一個隨機產生數字的Function時，我們可以
  * 在中間建立一個Target的介面讓adapter物件繼承，並在adapter中實作字串正數字的轉換
  * Client使用Target介面，來間接取得adapter物件，並使用adapter的轉換function，使原來的程式能夠直接取得正確格式的隨機數字並加以比較
![image](https://github.com/Lornzo/DesignPattern/blob/main/Adapter/images/example.png)