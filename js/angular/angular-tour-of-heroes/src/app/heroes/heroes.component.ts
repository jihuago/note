import { Component, OnInit } from '@angular/core';
import {Hero} from "../hero";

@Component({
  // 组件的选择器: 用来在父组件的模板中的匹配 HTML 元素的名称，以识别除该组件
  selector: 'app-heroes',
  // 组件模板文件的位置
  templateUrl: './heroes.component.html',
  // 组件私有 CSS 文件的位置
  styleUrls: ['./heroes.component.css']
})
export class HeroesComponent implements OnInit {

  // 添加属性
  // hero = 'Windstorm';
  hero: Hero = {
    id: 1,
    name: 'Windstorm'
  };

  constructor() { }

  // ngOnInit() 是一个生命周期钩子，Angular 在创建完组件后很快就会调用 ngOnInit()
  // ngOnInit() 一般放置初始化逻辑的地方
  ngOnInit(): void {
  }

}
