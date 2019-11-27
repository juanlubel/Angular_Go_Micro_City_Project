import {Component, EventEmitter, Host, Input, OnInit, Output} from '@angular/core';
import {UsersComponent} from '../users.component';

@Component({
  selector: 'app-users-list',
  templateUrl: './users-list.component.html',
  styleUrls: ['./users-list.component.css']
})
export class UsersListComponent implements OnInit {

  constructor(@Host() private Home: UsersComponent) { }
  // tslint:disable-next-line:no-input-rename
  @Input('dataSon') sonArray: string[];
  @Output() Select = new EventEmitter<string>();

  ngOnInit() {
    this.Select.emit('Hello Emit User');

    this.Home.hostedString = 'una string';
  }


}
