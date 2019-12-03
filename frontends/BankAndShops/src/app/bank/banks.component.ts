import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Banks, BankService } from '../core';
import { FormBuilder, FormGroup, Validators } from  '@angular/forms';

@Component({
  selector: 'app-banks',
  templateUrl: `banks.html`,
  styleUrls: ['bank.css']
})
export class BanksComponent implements OnInit{
  banks: Banks[];
  bankForm: FormGroup;

  constructor(
    private route: ActivatedRoute,
    private formBuilder: FormBuilder,
    private _http: BankService){
  }

  ngOnInit(){
    this.route.data.subscribe(
      (data: { banks: Banks[] }) => {
        this.banks = data.banks;
      }) 
      console.log(this.banks)
      this.bankForm  =  this.formBuilder.group({
     /*    BankName: ['', Validators.required] */
     Bank: ['', Validators.required],
     AccountNumber: ['', Validators.required],
     Owner: ['', Validators.required]
    });
  }
  createBank(){
    console.log(this.bankForm.value);
    this._http.addBank(this.bankForm.value).subscribe(
      result => {console.log(result); this.banks.push(result["bank"]); console.log(this.banks);})
  }
  login_bank(){
    console.log(this.bankForm.value);
    this._http.login_bank(this.bankForm.value).subscribe(
      result => {console.log(result);})
  }


}