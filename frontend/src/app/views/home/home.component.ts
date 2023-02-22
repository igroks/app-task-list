import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { ItemProps } from './home.model';
import { HomeService } from './home.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})

export class HomeComponent implements OnInit{
  listItemMap = {
    db1: [],
    db2: [],
  };
  items: ItemProps[] = [];
  nameControl = new FormControl('',[Validators.required, Validators.minLength(1)]);
  databaseForm!: FormGroup;
  formSubmitedd = false;

  constructor(
    private homeService: HomeService,
    private fb: FormBuilder
  ){ }

  ngOnInit(): void {
    this.loadItem('db1');
    this.loadItem('db2');
    this.databaseForm = this.fb.group({
      db1: false,
      db2: false
    })
  }

  reloadItems(){
    this.items = this.listItemMap.db1.concat(this.listItemMap.db2);
  }

  loadItem(db: string){
    this.homeService.getItems(db).subscribe((res) => {
      if(!!res) (this.listItemMap as any)[db] = res;
      this.reloadItems();
    });
  }

  isDatabasesValid(){
    return !(this.databaseForm.controls['db1'].value || this.databaseForm.controls['db2'].value) && this.formSubmitedd;
  }

  isNameValid() {
    return !this.nameControl.valid && this.formSubmitedd;
  }

  createItem(){
    this.formSubmitedd = true;
    if(this.nameControl.valid && this.databaseForm.valid){
      let item: ItemProps = {
        name: `${this.nameControl.value}`,
        createdAt: new Date()
      };
      let databaseMap = this.databaseForm.value;
      let databases = Object.keys(databaseMap).filter((database) => databaseMap[database]);

      databases.forEach((database) => this.homeService.insertItem(item, database).subscribe(
        () => {
          console.log('inserido com sucesso');
          this.loadItem(database);
        },
        () => {
          console.log('deu ruin');
        }
      ));
    } else {
      this.nameControl.updateValueAndValidity();
      this.databaseForm.updateValueAndValidity();
    }
  }
}
