import { Component, OnInit } from '@angular/core';
import {SocketService} from '../core/services';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
  providers: [SocketService]
})
export class LoginComponent {
}
