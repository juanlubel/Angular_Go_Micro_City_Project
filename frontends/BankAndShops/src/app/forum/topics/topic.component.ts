import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import { ActivatedRoute,Router } from '@angular/router';
import { FormBuilder, FormGroup, Validators } from  '@angular/forms';
import { MatBottomSheet, MatBottomSheetRef, MatDialogConfig, MatDialog } from '@angular/material';
import { CourseDialogComponent } from 'src/app/public/components/dialog.component';
import { Topic } from 'src/app/core';

@Component({
  selector: 'app-topics',
  templateUrl: `topics.html`,
  styleUrls: ['topics.css'],
})
export class TopicComponent implements OnInit{
  @Input("topics") topics:any
  @Output() ToggleCreate = new EventEmitter<any>();
  @Output() ToggleDelete = new EventEmitter<string>();
  constructor(
    private router: Router,
    private route: ActivatedRoute,
    private formBuilder: FormBuilder,
    private dialog: MatDialog) {}

  messages = []
  ngOnInit(){
    /* console.log(this.topics) */
    this.messages=this.topics
    console.log(this.messages)
  }
  goTopic(topic:string){
   this.router.navigate(['/forum/topic/', topic])
  }
  openDialog() {
    const dialogConfig = new MatDialogConfig();

  /*   dialogConfig.disableClose = true;
    dialogConfig.autoFocus = true; */

    dialogConfig.data = {
        id: 1,
        title: 'Create a topic',
        name: "Topic Tittle"
    };

    /* this.dialog.open(CourseDialogComponent, dialogConfig); */
    
    const dialogRef = this.dialog.open(CourseDialogComponent, dialogConfig);

    dialogRef.afterClosed().subscribe(
        data => {
          if(data && data.to_send!=null){
            /* console.log("Dialog output:", data); */
            this.onToggleCreate(data.to_send)
          }
        }
    );    
  }

  onToggleDelete(topic:string){
    this.messages = this.messages.filter((comment)=>comment.TopicTitle!=topic)
    this.ToggleDelete.emit(topic);
  }
  onToggleCreate(text:Topic){
    this.messages.push({Author:"UserTest",TopicTitle:text})
    this.ToggleCreate.emit({Author:"UserTest",TopicTitle:text});
  }
}
