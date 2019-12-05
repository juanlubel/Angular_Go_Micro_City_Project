import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { FormBuilder, FormGroup, Validators } from  '@angular/forms';

@Component({
  selector: 'app-forum',
  templateUrl: `forum.html`,
  styleUrls: ['forum.css']
})
export class ForumComponent implements OnInit{

  constructor(
    private route: ActivatedRoute,
    private formBuilder: FormBuilder,
    ){
  }
  messages = [
    {
      from: 'Entity 1',
      subject: 'Message Subject 1',
      content: 'Message Content 1'
    },
    {
      from: 'Entity 2',
      subject: 'Message Subject 2',
      content: 'Message Content 2'
    },
  ]
  ngOnInit(){

  }
 


}