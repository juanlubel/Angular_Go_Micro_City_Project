import {Component, OnInit} from '@angular/core';
import {UsersService} from '../services/users.service';

@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit {

  title = 'Hello Users';
  dataParent: string[] = ['uno', 'dos', 'tres'];
  hostedString: string;

  constructor(private usersService: UsersService) {
  }

  ngOnInit() {
    this.usersService.getAll()
      .subscribe((res: any) => {
      console.log(res);
    });
  }

  eventListener(event) {
    console.log(event);
  }
}
