# Facade (外觀模式/門面模式)
為子系統中的一組介面定義一個更高層且統一的介面，使子系統更容易使用

## 理解
為特定的目的，把多個分散在不同類別或不同系統(或Package)的功能，整合在一個新的介面裡面，以降低達成目的的程式複雜度

## 定義
* 意圖：簡化原有系統的使用方式(需要定義自已的介面)
* 問題：當我們只需要使用某個複雜系統的子集或是要以一種特殊的方式與系統溝通
* 方案：為原有的系統使用者提供一個量身訂做或是更容易使用的介面
* 角色：為使用者提供一個簡化的判面，使系統更容易使用
* 效果：Facade模式簡化了對所需子系統的使用過程，但是，由於Facade只使用了底層系統的「部份功能」，因此Facade的使用者會無法使用底層系統的某些功能
* 實作：
  * 定義一個 or 多個具備所需介面的新類別
  * 讓新的類別使用原有的系統

## 變形
* 在Facade類別下增加新的副程式，可以用以補充原有的功能，但是以不影響「簡化」程式使用為前題
* 用來隱藏 or 封裝系統(用Facade當作封裝層)
  * 用以追蹤係統的使用情況
  * 用以應付切換系統的需求

## Facade模式UML
![image](https://github.com/Lornzo/DesignPattern/blob/main/Facade/Facade.png)

## Facade模式對比
* 現在有兩個系統Client會用到「DB」、「Model」、「Element」，一般的情況下，兩個Client在Coding會各自使用到不同的物件功能，這樣程式碼會顯得很複雜  
![image](https://github.com/Lornzo/DesignPattern/blob/main/Facade/normal-pattern.png)

* 如果我們在中間使用一個Facade物件，統一整合後面的功能，在使用上就會顯得簡單一些
![image](https://github.com/Lornzo/DesignPattern/blob/main/Facade/facade-pattern.png)