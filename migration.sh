RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color
commands=$1

if [ "$commands" = "" ]
then
  echo "${RED}===========================${NC}"
  echo "${RED}PLEASE SPECIFY UP OR DOWN${NC}"
  echo "${RED}===========================${NC}"
fi

if [ "$commands" = "up" ]
then
  echo "${GREEN}===========================${NC}"
  echo "${GREEN}RUNNING MIGRATION${NC}"
  echo "${GREEN}===========================${NC}"
  (env $(cat .env | xargs) bash './migrate-up.sh')
fi

if [ "$commands" = "version" ]
then
  echo "${GREEN}===========================${NC}"
  echo "${GREEN}CHECK MIGRATION VERSION${NC}"
  echo "${GREEN}===========================${NC}"
  (env $(cat .env | xargs) bash './migrate-version.sh')
fi

if [ "$commands" = "down" ]
then
  echo "${GREEN}===========================${NC}"
  echo "${GREEN}REVERT MIGRATION${NC}"
  echo "${GREEN}===========================${NC}"
  (env $(cat .env | xargs) bash './migrate-down.sh')
fi