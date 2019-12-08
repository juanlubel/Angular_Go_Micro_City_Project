import {MAT_DIALOG_DATA, MatDialogRef} from "@angular/material";
import { Component, OnInit, Inject } from '@angular/core';
import { FormGroup, FormBuilder } from '@angular/forms';

@Component({
    selector: 'course-dialog',
    templateUrl: './test.html',
    /* styleUrls: ['./test.css'] */
})
export class CourseDialogComponent implements OnInit {

    form: FormGroup;
    to_send:string;
    data:{}

    constructor(
        private fb: FormBuilder,
        private dialogRef: MatDialogRef<CourseDialogComponent>,
        @Inject(MAT_DIALOG_DATA) data) {
        this.to_send = data.to_send;
        this.data=data
    }

    ngOnInit() {
        this.form = this.fb.group({
            to_send: [this.to_send, []],
        });
    }

    save() {
        this.dialogRef.close(this.form.value);
    }

    close() {
        this.dialogRef.close();
    }
}