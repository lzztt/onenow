FROM jekyll/minimal:4.2.0 AS build
WORKDIR /tmp
COPY . .
RUN ./build.sh

FROM jekyll/minimal:4.2.0
WORKDIR /tmp
COPY --from=build /tmp/_site ./_site
RUN echo -e 'skip_initial_build: true\nlivereload: false' > _config.yml
CMD jekyll serve --no-watch