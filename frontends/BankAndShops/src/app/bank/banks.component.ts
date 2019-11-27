import { Component, Directive, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Banks } from '../core';
import { FormControl, FormGroup, FormBuilder } from '@angular/forms';

@Component({
  selector: 'app-banks',
  templateUrl: `banks.html`,
  styles: []
})
export class BanksComponent implements OnInit{
  banks: Banks;
  contactForm: FormGroup

  // Step 1
  createFormGroup() {
    return new FormGroup({
      personalData: new FormGroup({
        email: new FormControl(),
        mobile: new FormControl(),
        country: new FormControl(),
      }),
      requestType: new FormControl(),
      text: new FormControl(),
    })
  }
  constructor(private route: ActivatedRoute,private formBuilder: FormBuilder) {
    this.contactForm = this.createFormGroup();
  }

  ngOnInit(){
    this.route.data.subscribe(
      (data: { banks: Banks }) => {
        this.banks = data.banks;
      }) 
      console.log(this.banks)
  }
  


}