import {Component, OnInit} from '@angular/core';
import {LoginService} from '../core/services';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  private token: any;

  constructor(
    private userService: LoginService
  ) { }

  ngOnInit() {
  }

  login(a: string, b: string) {
    console.log(a, b);
    const params = {
      name: a,
      pass: b
    };
    this.userService.login(params).subscribe(res => {
      console.log(res, 'error');
      this.token = res.token;
    });
  }
}
