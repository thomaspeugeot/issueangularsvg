import { ComponentFixture, TestBed } from '@angular/core/testing';

import { IssueangularsvgspecificComponent } from './issueangularsvgspecific.component';

describe('IssueangularsvgspecificComponent', () => {
  let component: IssueangularsvgspecificComponent;
  let fixture: ComponentFixture<IssueangularsvgspecificComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [IssueangularsvgspecificComponent]
    });
    fixture = TestBed.createComponent(IssueangularsvgspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
