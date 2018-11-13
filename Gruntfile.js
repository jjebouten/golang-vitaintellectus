/*global module:false*/
module.exports = function (grunt) {
    // Project configuration.
    grunt.initConfig({
            uglify: {
                options: {
                    sourceMap: true
                },
                main: {
                    src: [
                        'Private/Scripts/main.js'
                    ],
                    dest: 'Public/Scripts/main.min.js'
                },
            },
            less: {
                main: {
                    options: {
                        paths: ["Less"],
                        compress: true,
                        plugins: [
                            new (require('less-plugin-autoprefix'))({browsers: ["> 1%", "last 2 versions", "Firefox ESR", "Opera 12.1"]}),
                        ]
                    },
                    src: ['Private/Less/main.less'],
                    dest: 'Public/Css/main.css'

                }
            },
            watch: {
                less: {
                    files: ['Private/Less/**/*.less'],
                    tasks: ['less'],
                    options: {
                        // Start a live reload server on the default port 35729
                        livereload: true,
                        // No live reload on compile error.
                        livereloadOnError: false
                    }
                },
                javascript: {
                    files: ['Private/Scripts/*.js'],
                    tasks: ['javascript-main'],
                    options: {
                        // Start another live reload server on port 35729
                        livereload: true,
                        // No live reload on compile error.
                        livereloadOnError: false
                    }
                }

            }
        }
    );

    // These plugins provide necessary tasks.
    grunt.loadNpmTasks('grunt-contrib-uglify');
    grunt.loadNpmTasks('grunt-contrib-less');
    grunt.loadNpmTasks('grunt-contrib-watch');

    grunt.registerTask('js-main', 'Concatenates and minify main javascript file', ['uglify:main']);

    // deprecated/legacy
    grunt.registerTask('javascript-main', 'Concatenates and minify main javascript file', ['uglify:main']);
    grunt.registerTask('javascript-compress', 'Concatenates and minify vendor javascript file', ['uglify:compress']);
    grunt.registerTask('watch', 'Watches less files and main.js', ['watch']);

    // By default run all
    grunt.registerTask('default', ['uglify', 'less']);
};
