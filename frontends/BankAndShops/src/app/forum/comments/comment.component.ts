import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from  '@angular/forms';

export interface Message {
  text: string;
  name: string;
}
@Component({
  selector: 'app-comment',
  templateUrl: './comment.html',
  styleUrls: ['./comment.component.css'],
})
export class CommentComponent implements OnInit {
    constructor(private http: HttpClient,private formBuilder: FormBuilder,) {}
    /* @Output() onSendMessage: EventEmitter<Message> = new EventEmitter();
    message = {
      name: '',
      text: '',
    }; */
    comment_area: FormGroup;
    sendMessage() {
      console.log("hola")
    }
    ngOnInit() {
      this.comment_area =  this.formBuilder.group({
        /*    BankName: ['', Validators.required] */
        comment_user: ['', Validators.required],
        comment_message: ['', Validators.required]
       });
    }
  }