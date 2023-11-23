# Visual issue when mixing angular material component and angular svg component

This issue has been reproduced on windows/mac os with FX / Chrome / Edge browsers.

## Tools tools for bug reproduction

nNeeded tools are : 

- git
- npm
- go (>=1.21) for the backend.

## Sequence

```bash
git clone https://github.com/thomaspeugeot/issueangularsvg.git
cd issueangularsvg/ng
npm i
ng build
cd ../go/cmd/issueangularsvg

```