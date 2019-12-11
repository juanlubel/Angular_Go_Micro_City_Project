import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { FormBuilder, FormGroup, Validators } from  '@angular/forms';
import { Topic,ForumService, Comment } from '../core';

@Component({
  selector: 'app-forum',
  templateUrl: `forum.html`,
  styleUrls: ['forum.css']
})
export class ForumComponent implements OnInit{

  topics:Topic[]

  constructor(
    private route: ActivatedRoute,
    private formBuilder: FormBuilder,
    private ForumService:ForumService
    ){
  }
  ngOnInit(){
    this.route.data.subscribe(data => {
        console.log(data)
        this.topics =data.topics
    }) 
  }
  createTopic(e:Topic){
    console.log(e)
    this.ForumService.createTopic(e).subscribe()
  }


}