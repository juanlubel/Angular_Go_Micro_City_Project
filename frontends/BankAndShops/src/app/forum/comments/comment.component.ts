import { Component, OnInit, Output, EventEmitter, Host } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from  '@angular/forms';
import { ActivatedRoute } from '@angular/router';
import { TopicWithComments, Comment,ForumService} from 'src/app/core';



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
    constructor(
      private ForumService:ForumService,
      private formBuilder: FormBuilder,
      private route: ActivatedRoute,
    ) {}
    topicWithComments:TopicWithComments
    toCreateComment:Comment
    comment_area: FormGroup;
    
    ngOnInit() {
      this.route.data.subscribe(data => {
        console.log(data)
        this.topicWithComments =data.topic;
      }) 
      this.comment_area =  this.formBuilder.group({
        comment_message: ['', Validators.required]
       });
    }
    sendMessage() {
      console.log(this.comment_area.value.comment_message)
      this.toCreateComment = {
        Author:"test",
        Body:this.comment_area.value.comment_message,
        TopicTittle:this.topicWithComments.Topic.TopicTitle}
      this.topicWithComments.Comments.push(this.toCreateComment)
     this.ForumService.createComment(this.toCreateComment,"").subscribe()
    }
  }