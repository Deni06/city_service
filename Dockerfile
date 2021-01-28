FROM alpine
ADD CityService-srv /CityService-srv
ENTRYPOINT [ "/CityService-srv" ]
