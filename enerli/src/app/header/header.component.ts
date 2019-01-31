import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit {
  public headerData = {
    shortName: 'SV',
    time:'17:56',
    date: '30th november 2018'
  };

  constructor() { }

  ngOnInit() {
  }

  alertMessage() {
    alert('open user menu');
  }

}
