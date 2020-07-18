const gulp = require('gulp');
const sass = require('gulp-sass');

const pre_path = './react/';

const sass_SOURCE = pre_path + 'src/scss/**/*';
const sass_THEME = pre_path + 'src/scss/theme.scss';
const sass_DISTRIBUTION = pre_path + 'src/css';

gulp.task('copy-sass', () => {
  return gulp.src(sass_THEME)
    .pipe(sass())
    .pipe(gulp.dest(sass_DISTRIBUTION));
});

gulp.task('watch-sass', () => {
  gulp.watch(sass_SOURCE, gulp.series('copy-sass'));
});