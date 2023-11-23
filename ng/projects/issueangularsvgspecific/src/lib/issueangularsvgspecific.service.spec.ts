import { TestBed } from '@angular/core/testing';

import { IssueangularsvgspecificService } from './issueangularsvgspecific.service';

describe('IssueangularsvgspecificService', () => {
  let service: IssueangularsvgspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(IssueangularsvgspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
