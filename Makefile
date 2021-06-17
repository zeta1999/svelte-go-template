setupfront:
	cd frontend && npm install

setupback:
	cd backend && go get

devback:
	cd backend && air

devfront:
	cd frontend && npm run dev

