module.exports = function(grunt) {
    grunt.initConfig({
        pkg: grunt.file.readJSON('package.json'),
        // Deletes the build-folder and its contents
        clean: {
            build: {
                src: ['build']
            }
        },
        // Compresses the build-folder into zip or tar.gz
        compress: {
            zip: {
                options: {
                    archive: "build.zip"
                },
                expand: true,
                cwd: 'build/',
                src: ['**/*'],
                dest: '.'
            },
            tar: {
                options: {
                    archive: "build.tar.gz"
                },
                expand: true,
                cwd: 'build/',
                src: ['**/*'],
                dest: '.'
            }
        },
        // Copies everything except less-related files to build folder
        copy: {
            build: {
                cwd: 'src',
                src: ['**', '!**/*.less', '!**/less*.js'],
                dest: 'build',
                expand: true
            }
        },
        // Compiles less to CSS
        less: {
            build: {
                files: [{
                    expand: true,
                    cwd: 'src',
                    src: ['**/*.less'],
                    dest: 'build',
                    ext: '.css'
                }]
            }
        },
        replace: {
            // Replaces less-related content in php-files with css paths
            less: {
                src: ['build/**/*.html'],
                overwrite: true,
                replacements: [
                    {
                        from: '<script src="js/less-1.7.0.min.js"></script>'
                    },
                    {
                        from: /<link rel="stylesheet\/less\" type=\"text\/css\" href=\"(css\/styles).less\" \/>/g,
                        to: '<link rel="stylesheet" href="$1.css" />'
                    }
                ]
            },
            // Replace console.log
            consolelog: {
                src: ['build/**/*.js'],
                overwrite: true,
                replacements:[
                    {
                        from: /console\.log\(.*\);/g
                    }
                ]
            }
        }
    });

    // Load tasks
    grunt.loadNpmTasks('grunt-contrib-clean');
    grunt.loadNpmTasks('grunt-contrib-compress');
    grunt.loadNpmTasks('grunt-contrib-copy');
    grunt.loadNpmTasks('grunt-contrib-less');
    grunt.loadNpmTasks('grunt-text-replace');

    // A very basic default task.
    grunt.registerTask('default', 'Log some stuff.', function() {
        grunt.log.write('\n');
        grunt.log.ok('Use argument build for production\n');
    });

    // Runs all tasks needed for a build
    grunt.registerTask('build', 'Buildtask', ['clean', 'copy', 'less', 'replace:less', 'replace:consolelog']);
};