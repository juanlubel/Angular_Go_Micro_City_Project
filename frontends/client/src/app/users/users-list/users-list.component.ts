import {Component, EventEmitter, Host, Input, OnInit, Output} from '@angular/core';
import {UsersComponent} from '../users.component';
import {User} from '../../models';

@Component({
  selector: 'app-users-list',
  templateUrl: './users-list.component.html',
  styleUrls: ['./users-list.component.css']
})
export class UsersListComponent implements OnInit {

  constructor(@Host() private Home: UsersComponent) { }
  // tslint:disable-next-line:no-input-rename
  @Input('dataSon') sonArray: User[];
  @Output() Select = new EventEmitter<string>();

  ngOnInit() {
    this.Select.emit('Hello Emit User');
    console.log(this.sonArray);
    this.Home.hostedString = 'una string';
  }


}
