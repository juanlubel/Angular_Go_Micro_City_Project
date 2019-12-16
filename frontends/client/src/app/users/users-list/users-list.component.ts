import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {User} from '../../models';

@Component({
  selector: 'app-users-list',
  templateUrl: './users-list.component.html',
  styleUrls: ['./users-list.component.css']
})
export class UsersListComponent implements OnInit {
  // tslint:disable-next-line:no-input-rename
  @Input('users') users: User[];
  @Output() Select = new EventEmitter<string>();
  constructor() {

  }
  ngOnInit() {

  }

  selectUser(slug: string) {
    console.log(slug);
    this.Select.emit(slug);
  }

}
